my $path = ".";

sub scan_file {
    my @files = glob(@_[0]);
    foreach(@files) {
        if(-d $_) { # dir 
            my $path = "$_/*";
            scan_file($path);
        }
        elsif(-f $_) { # file
            # if($_ =~ /log$/) {
            #     print "$_\n";
            # } elsif ($_ =~ /yaml$|yml$/) {
            #   print "$_\n";
            # }
            if ($_ =~ /.xx$/) {
                # print "$_\n";
                replace($_);
            }
        }
    }
}

sub replace {
    $file = glob(@_[0]);
    $srcfile = $file;
    $dstfile = sprintf("%s.copy", $file);

    open(outStream, sprintf("> %s", $dstfile)) || die "$!\n";
    open(inStream, sprintf("< %s", $srcfile)) || die "cannot open the stream: $!\n";

    # $packageName = "server";
    $srcModuleName = "truffle";
    $dstModuleName = "demo";

    while (<inStream>) {
        if ($_ =~ m/^package /) {
            print outStream sprintf("package %s", $'); # do nothing
        } elsif ($_ =~ s/$srcModuleName\//$dstModuleName\//) {
            print outStream $_;
        } else {
            print outStream $_;
        }
    }
    close inStream;
    close outStream;
    unlink $srcfile;
    rename ($dstfile, $srcfile);
}

scan_file($path);
